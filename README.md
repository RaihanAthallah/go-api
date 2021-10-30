# Golang Docker Compose with Private Repository

Sample how to create docker for golang with Private Repository

## Environments
- Development (Hot Reload)
- Production

## Setup
Open and edit docker-compose files:
```
- repo_user={YOURUSERNAME}
- repo_token={GENERATEDTOKEN}
- repo_url={PRIVATEREPO} # git.mydomain.com
- repo_organization={YOURORGANIZATION} # useraccount/organization ex: git.mydomain.com/organization/repo
- app_name={APPNAME} # appname 
```

## How To Generate Token User
### Gitea
-  Goto <b> Profile Setting</b>
-  Tab <b>Application</b>
-  On Section <b>Generate Token</b> fill <i>Token Name</i> then Generate Token
![image](https://user-images.githubusercontent.com/10952676/125617887-913e306e-8de5-469d-91d8-852dc569a867.png)

### Github
- Visit the **Settings** of the user account.
- Goto **Developer Settings**
- Then choose **Personal Access Token**
![image](https://user-images.githubusercontent.com/10952676/125618327-5e26d42b-b26d-47c2-b1bc-cf38fc41a150.png)


## Usage
- Run as Development

    ```
    docker-compose -f docker-compose.dev.yml up --build
    ```
- Run as Production
  
    ```bash
    docker-compose -f docker-compose.yml up --build
    ```