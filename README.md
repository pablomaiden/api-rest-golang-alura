# api-rest-golang-alura
Este projeto é basicamente para estudos e aprofundamento dos conhecimentos na linguagem GO

Para instalar o gorila mux
go get -u github.com/gorilla/mux


Existe um comando que conseguimos acessar essa máquina e descobrir qual é o host dela, podemos fazer assim: docker-compose exec, o nome do serviço que queremos subir que é esse aqui, queremos visualizar qual é a porta desse serviço aqui, do serviço do postgres, coloco postgres sh. docker-compose exec postgres sh.

[01:16] Dou um "Enter" e eu estou nessa máquina do postgres com sh. Para eu ver qual é o caminho dela eu posso colocar aqui # hostname -i e quando eu dou um "Enter", olha que interessante, ele vai me dá um endereço “172.19.0.2”. Esse valor aqui é o valor que eu preciso passar aqui dentro para conseguir com que o pgAdmin e aquele banco de dados postgres que temos conversem.

[01:38] Outra forma que temos para conseguir pegar esse valor é assim, eu digito: docker inspect, passo qual é o nome desse meu banco de dados, vamos só colocar aqui: docker ps, vou rodar aqui e temos o banco de dados do postgres que é "461", esse é o ID do container que queremos.

[02:28] Eu vou colocar esse comando docker inspect 461| grep IPAddress e ele me dá aqui a mesma informação: "172.19.0.2". Temos duas formas, ou acessamos a máquina com exec do postgres sh e pergunta qual é o host name ou rodamos esse comando aqui. Esse comando precisamos saber só qual é o ID do container que estamos rodando o postgres.
