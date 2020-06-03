## 테스트 환경
- OS: Ubuntu 18.04 (raspberry)
* * *

### 1. 프로젝트 빌드 & 데몬으로 시작
- 리눅스의 경우는 `sudo docker`로 무조건 시작해야 합니다.<br><br>
- 도커 프로젝트 빌드:<br>`docker-compose build --parallel` or `docker-compose build`<br><br>
- 도커 프로젝트 시작:<br>`docker-compose up -d`<br><br>


### 2. UYeG 등록하기
- localhost:8000/admin로 접속하여 UYeG 등록.

### 3. Tips
 - `docker logs <container>`로 로그 확인 가능
 - `docker stats`로 컨테이너별 리소스 사용 현황 확인 가능<br>[코어 할당 갯수에 따라 CPU 최대 %가 다름 (코어 갯수 * 100%)]
 - `docker-compose.yml` 파일에서 web의 환경변수로 admin 유저의 초기 이메일 패스워드를 설정 가능.
