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

# Desafio 03 - Pipeline de CI com SonarCloud
```
Crie uma pequena aplicação simples em node.js (qualquer aplicação mesmo)
[R] Código: integracao-continua/projeto-demo/math.go
Crie testes de unidade para essa aplicação
[R] Código: integracao-continua/projeto-demo/math_test.go 
Crie uma pipeline de CI utilizando o Github actions que:
- Instale a aplicação 
- Execute os testes
- Execute o SonarCloud
[R] .github/workflows/ci.yaml
```

## Anotações
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
