package usecase

import (
	"crypto/md5"
	"encoding/hex"

	entity "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
	response "github.com/hayvee-website-development/go-api-hayvee/app/model/response/doctor"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
	logger "github.com/hayvee-website-development/go-api-hayvee/infrastructure/io"
	"github.com/jinzhu/copier"
	"github.com/thanhpk/randstr"
)

type userAccessUsecase struct {
	BaseRepository       repository.BaseRepository
	UserAccessRepository repository.UserAccessRepository
	UserRepository       repository.UserRepository
}

type UserAccessUsecase interface {
	GenerateToken(input entity.HvUserAccess) string
	GenerateHashPasswordForLogin(iduser int, password string) string
	GenerateHashPasswordForRegister(salt string, password string) string
	ValidTokenWithID(id int, token string) bool
}

func NewUserAccessUsecase(
	br repository.BaseRepository,
	uar repository.UserAccessRepository,
	sar repository.UserRepository,
) UserAccessUsecase {
	return &userAccessUsecase{br, uar, sar}
}

func (uu *userAccessUsecase) GenerateToken(input entity.HvUserAccess) string {
	cl := logger.WithFields(logger.Fields{"UserInteractor": "Token"})
	var token string
	cekAcc, _ := uu.UserAccessRepository.Find(input.IDUser)
	if cekAcc.Token == "" {
		token = randstr.String(32)
		input.Token = token
		err := uu.UserAccessRepository.Create(input)
		if err != nil {
			cl.Errorf("[ERROR]  %v", err.Error())
		}
	} else {
		token = randstr.String(32)
		input.Token = token
		err := uu.UserAccessRepository.Update(input.IDUser, token)
		if err != nil {
			cl.Errorf("[ERROR]  %v", err.Error())
		}
	}
	return token
}
func (uu *userAccessUsecase) GenerateHashPasswordForLogin(iduser int, password string) string {
	findusersuperadmin, _ := uu.UserRepository.Find(iduser)
	responseusersuperadmin := response.ResultUser{}
	copier.Copy(&responseusersuperadmin, &findusersuperadmin)

	passwordplussalt := password + responseusersuperadmin.Salt
	passwordmd5 := md5.Sum([]byte(passwordplussalt))
	md5tostring := hex.EncodeToString(passwordmd5[:])
	passwordmd5plussalt := md5tostring + responseusersuperadmin.Salt
	passwordmd5final := md5.Sum([]byte(passwordmd5plussalt))
	md5tostringfinal := hex.EncodeToString(passwordmd5final[:])

	return md5tostringfinal
}
func (uu *userAccessUsecase) GenerateHashPasswordForRegister(salt string, password string) string {

	passwordplussalt := password + salt
	passwordmd5 := md5.Sum([]byte(passwordplussalt))
	md5tostring := hex.EncodeToString(passwordmd5[:])
	passwordmd5plussalt := md5tostring + salt
	passwordmd5final := md5.Sum([]byte(passwordmd5plussalt))
	md5tostringfinal := hex.EncodeToString(passwordmd5final[:])

	return md5tostringfinal
}

func (uu *userAccessUsecase) ValidTokenWithID(id int, token string) bool {
	res := uu.UserAccessRepository.ValidTokenWithID(id, token)
	return res
}
