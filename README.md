# Port-Container-Management-Server
항만 컨테이너의 관리를 위한 서비스


## 지원 기술
* 항만컨테이너 자동통합검색플랫폼을 위한 라우팅 코어 기술 
    - 컨테이너 사전 정보 및 컨테이너 검색 결과 관리용 API 
    - 항만 검색기 운영정보 및 상태 정보 관리용 API
* 컨테이너 및 검색기 관리를 위한 Database
* 컨테이너 및 검색기 정보 시각화를 위한 React 웹


## Start
```
1. $ git clone https://github.com/jinsung-cho/Port-Container-Management-Server.git
```
```
2. $ cd Port-Container-Management-Server
```
```
3. $ mv .env-sample .env
```

4. Edit the .env file
```
# api-server
API_SERVER_HOST=10.0.0.1    # API_SERVER가 실행되는 HOST IP
API_SERVER_PORT=10000       # API_SERVER가 실행되는 HOST의 PORT
DB_SERVER_HOST=10.0.0.1     # API_SERVER와 연결될 DB_SERVER의 HOST IP
DB_SERVER_PORT=9999         # API_SERVER와 연결될 DB_SERVER의 HOST의 PORT

# db-server
DB_SERVER_HOST=10.0.0.1     # DB_SERVER가 실행되는 HOST IP
DB_SERVER_PORT=9999         #

# postgres
POSTGRES_USER=test           #
POSTGRES_PASSWORD=test       #
POSTGRES_PORT=5432           #
POSTGRES_HOST=10.0.0.1       #
POSTGRES_DBNAME=Application  #

#postgrest                   #
POSTGREST_PORT=4000          #
POSTGREST_SECRET=qkkejklwvblkrenbklenbklernklvbenrklvnre
                             #
# TOS
TOS_HOST=10.10.10.10         #
TOS_PORT=2000                #
TOS_PATH=rest-path           #

# WEB
WEB_PORT=3001                #
NEXT_PUBLIC_API_BASE_URL=http://10.0.0.0:4000
NEXT_PUBLIC_API_TOKEN=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidGVzdCJ9.Ud1nm095AmrhaYFrgDPHvBlm3W7GQ7fz7xUQbG8aZ-U
```

### 작성 예
```
# api-server
API_SERVER_HOST=10.0.0.1    
API_SERVER_PORT=10000    
DB_SERVER_HOST=10.0.0.1   
DB_SERVER_PORT=9999  

# db-server
DB_SERVER_HOST=10.0.0.1 
DB_SERVER_PORT=9999      

# postgres
POSTGRES_USER=test           
POSTGRES_PASSWORD=test       
POSTGRES_PORT=5432           
POSTGRES_HOST=10.0.0.1       
POSTGRES_DBNAME=Application  

#postgrest                   
POSTGREST_PORT=4000          
POSTGREST_SECRET=qkkejklwvblkrenbklenbklernklvbenrklvnre
                             
# TOS
TOS_HOST=10.10.10.10        
TOS_PORT=2000                
TOS_PATH=rest-path           

# WEB
WEB_PORT=3001                
NEXT_PUBLIC_API_BASE_URL=http://10.0.0.0:4000
NEXT_PUBLIC_API_TOKEN=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidGVzdCJ9.Ud1nm095AmrhaYFrgDPHvBlm3W7GQ7fz7xUQbG8aZ-U

```




