# Golang Korea X AWSKRUG Codelab

2019.03.30 Golang Korea X AWSKRUG 코드랩 프로젝트

이 코드랩에서는 Go를 사용해 OAuth2를 통한 간단한 소셜 로그인을 구현하면서 Go 프로그래밍에 익숙해지는걸 목표로 합니다.  

## Project Setup

```console
$ git clone -b skeleton https://github.com/golangkorea/codelab
$ cd codelab
$ ./clean.sh
```

## Project Structure

```
├── config
├── handler
├── model
├── oauth
├── static
│   ├── images
│   └── stylesheets
├── templates
├── main.go
└── config.yaml
```

- **config**: 설정값 관리 패키지
- **handler**: 핸들러 (컨트롤러) 관리 패키지
- **model**: 모델 구조체 패키지 
- **oauth**: OAuth2 인증 관련 패키지
- **static**: 정적 파일 모음 (스타일시트, 이미지등)
- **templates**: 템플릿 파일 모음
