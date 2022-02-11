package usecase

import (
	"strconv"
	"time"

	entity "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
	req "github.com/hayvee-website-development/go-api-hayvee/app/model/request/user"
	response "github.com/hayvee-website-development/go-api-hayvee/app/model/response/doctor"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
	"github.com/hayvee-website-development/go-api-hayvee/config"
	"github.com/hayvee-website-development/go-api-hayvee/pkg/helper"
	"github.com/jinzhu/copier"
	"github.com/thanhpk/randstr"
)

type userUsercase struct {
	BaseRepository       repository.BaseRepository
	UserRepository       repository.UserRepository
	UserAccessRepository repository.UserAccessRepository
	UserDataRepository   repository.UserDataRepository
	UserAccessUsecase    UserAccessUsecase
}

type UserUsercase interface {
	Find(id int, token string) (interface{}, error)
	FindAll(token string, name string, email string, phone string, id_identifier int) (result []entity.HvUserData, err error)
	VerifyEmail(username string, password string) interface{}
	IsDuplicateEmail(username string) bool
	Create(user req.RegRegisterUser) (interface{}, error)
	CreateIdentity(identity req.RegRegisterUserIdentity, token string) (interface{}, error)
}

func NewUserUsercase(
	br repository.BaseRepository,
	sar repository.UserRepository,
	saar repository.UserAccessRepository,
	sadr repository.UserDataRepository,
	saa UserAccessUsecase,
) UserUsercase {
	return &userUsercase{
		br,
		sar,
		saar,
		sadr,
		saa,
	}
}

func (uu *userUsercase) Find(id int, token string) (interface{}, error) {
	var err error
	findUser, err := uu.UserRepository.Find(id)
	res := response.ResultUser{}
	copier.Copy(&res, &findUser)

	resp := response.ResponeRegister{}
	resp.JwtToken, _ = helper.GenerateJwt(strconv.Itoa(res.ID), res.Email,
		config.C.Auth.ApplicationIssuer, config.C.Auth.CmsSecret, config.C.Auth.ExpiredTimeHour)
	resp.User = res
	resp.Token = token
	resp.ServerTime = time.Now().Format("2006-01-02 15:04:05")

	return resp, err
}

func (uu *userUsercase) FindAll(token string, name string, email string, phone string, id_identifier int) (result []entity.HvUserData, err error) {
	validasitoken := uu.UserAccessRepository.ValidToken(token)

	Filtersuperadmindata := map[string]interface{}{
		"name":          name,
		"email":         email,
		"phone":         phone,
		"id_identifier": id_identifier,
	}

	if validasitoken {
		filterlistdoctor, _ := uu.UserDataRepository.FindByParam(Filtersuperadmindata)
		if filterlistdoctor != nil {
			return uu.UserDataRepository.FindByParam(Filtersuperadmindata)
		}
		return uu.UserDataRepository.FindAll()
	}
	return nil, err
}

func (uu *userUsercase) VerifyEmail(email string, password string) interface{} {
	filter := map[string]interface{}{
		"email": email,
	}
	res, err := uu.UserRepository.FindByParam(filter)
	if err != nil {
		return false
	}
	uda := entity.HvUserAccess{
		IDUser:       res.IDUser,
		IDIdentifier: res.IDIdentifier,
	}
	token := uu.UserAccessUsecase.GenerateToken(uda)
	passwordhasher := uu.UserAccessUsecase.GenerateHashPasswordForLogin(res.IDUser, password)
	if res.Email == email && res.Password == passwordhasher {
		resp, _ := uu.Find(res.IDUser, token)
		return resp
	}
	return false
}

func (uu *userUsercase) IsDuplicateEmail(email string) bool {
	filter := map[string]interface{}{
		"email": email,
	}
	res, err := uu.UserRepository.FindByParam(filter)
	if err != nil {
		return true
	}
	if res.IDUser == 0 {
		return true
	}
	return false
}

func (uu *userUsercase) Create(input req.RegRegisterUser) (interface{}, error) {
	uu.BaseRepository.BeginTx()
	salt := randstr.String(32)
	passwordhashgenerator := uu.UserAccessUsecase.GenerateHashPasswordForRegister(salt, input.Password)
	superadmin := entity.HvUser{
		Email:        input.Email,
		Password:     passwordhashgenerator,
		IDIdentifier: 4,
		Salt:         salt,
	}
	rSuperAdmin, errrSuperAdmin := uu.UserRepository.Create(superadmin)
	if errrSuperAdmin != nil {
		uu.BaseRepository.RollbackTx()
		return false, errrSuperAdmin
	}
	superadmindata := entity.HvUserData{
		IDUser:       rSuperAdmin.IDUser,
		Email:        input.Email,
		Phone:        input.Handphone,
		Name:         input.Name,
		IDIdentifier: 4,
	}
	rProfileSuperAdmin, _ := uu.UserDataRepository.Create(superadmindata)

	superadminaccess := entity.HvUserAccess{
		IDUser:       rSuperAdmin.IDUser,
		IDIdentifier: 4,
	}

	token := uu.UserAccessUsecase.GenerateToken(superadminaccess)
	resp, err := uu.Find(rProfileSuperAdmin.IDUser, token)
	if err != nil {
		uu.BaseRepository.RollbackTx()
		return false, err
	}
	return resp, err
}

func (uu *userUsercase) CreateIdentity(identity req.RegRegisterUserIdentity, token string) (interface{}, error) {
	findiduser, _ := uu.UserAccessRepository.FindIDByToken(token)
	uu.BaseRepository.BeginTx()
	useridentity := entity.HvUserData{
		Umur:         identity.Umur,
		JenisKelamin: identity.JenisKelamin,
		Alamat:       identity.Alamat,
		Kota:         identity.Kota,
	}
	createidentity, err := uu.UserDataRepository.CreateIdentity(useridentity, findiduser.IDUser)
	if err != nil {
		uu.BaseRepository.RollbackTx()
		return false, err
	}
	return createidentity, err

}
