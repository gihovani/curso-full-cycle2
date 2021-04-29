# curso-full-cycle2
Repositório com os arquivos do Curso Full Cycle 02 da Code Education

# Desafio 01 - Desafio Go (Docker hub):
```
cd docker/01-desafio-go
docker run --rm gihovani/codeeducation
```

# Desafio 02 - Desafio Nginx com Node.js (docker-compose + node + nginx + mysql):
```
cd docker/02-desafio-nginx
chmod +x node/docker-entrypoint.sh 
docker-compose up -d  
```

# Executar SonarQube
```
docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest
```

# Gerar Arquivo coverage.out
```
cd ./integracao-continua/projeto-demo/
go test -coverprofile=coverage.out
```

# Executar SonarScanner
```
docker run \
    --rm \
    -v "$(pwd)/integracao-continua/projeto-demo/:/usr/src" \
    sonarsource/sonar-scanner-cli
```

# Aluno 
Gihovani Filipp <gihovani@gmail.com>
