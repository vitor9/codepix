#Cockburn, ou Dr. Alistair Cockburn, disse o seguinte
# Ports and Adapters

##Aplicação no meio 
que resolve problemas da aplicação

## Adaptadores em volta
para que possam se comunicar com ela.

Qualquer outro sistema pode se conectar na minha aplicação, atraveś de portas que irei disponibilizar. E irei ter adaptadores para a minha aplicação se conectar a algumas outras coisas ali também.



# Jeffrey Palermo, MVP, Trabalha na Clear Measure
## 2008 ele veio com a Onion Archtecture

O que é essa aplication?
# Aplication Core
## Domain Model
Meu dominio ou minha entidade. São as coisas únicas que irão definir a minha aplicação.

## Domain Services
Camada aonde esses models e essas coisas podem se falar ou executar transações entre si

## Aplication Services
Camadas aonde vão lidar com complexidade para fazer acesso a camada de dominio

## Por fora dessas camadas...
Temos testes, infraestrutura, user interface, infraestrutura


# Em 2000, Robert C. AKA Uncle BOB
## Um dos autores do manifesto agil
## Criador do SOLID em 2000
## Famoso pelo conceito de Clean Code em 2009
## Clean architecture em 2017
<img>

</img>

## Entities
Fica o que é o coração, as coisas unicas na minha aplicação, substantivos, como: codepix, banco, transacao, conta-bancarias.

## Use Cases
Ditam o fluxo de como as coisas funcionam, por exemplo: cria um banco, depois cria uma conta e depois gera uma transação. Todas essas regras, já estão na regra de entidade, mas o use case, gera o fluxo e acaba integrando essas outras camadas.

## Controladores
Integrações de pessoas chegando na aplicação

## Apresentações
Formatos e camadas prontas para presentações e trazer dados que importam

## Gateways
para trabalhar com integrações


Vamos trabalhar com uma variação desses 3 formatos, utilizando padrões de projeto muito concilidado e utilizados pelo DDD


#Nossa estrutura vai funcionar da seguinte forma:

# Application

## factory
gera objetos prontos que tem dependencia. Gerar uma conexão com banco de dados, gerar repositorio e botar banco dentro, gerar um usecase e botar o repositorio dentro. Vamos ter os objetos prontos e a factory vai fazer isso para a gente.

## grpc
servidor e serviços disponibilizados via gRPC. Tudo de gRPC ta ai, servidor, serviço.

## kafka
Consumo e processamento de transações com o Apache Kafka. Se eu quero consumir e processar informações usando kafka, vai ta tudo ai. Por que ai vai ter a complexidade de nos comunicarmos com outros micro-serviços. Esse tipo de comunicação é a complexidade da aplicação.
O dominio vai receber essas informações e vai consolidar.

## model
Estrutura de objetos que receberão as requisições externas(via Kafka ou gRPC). 
Ajuda a modelar os dados que estão vindo de fora.
Essa camada é diferente do que temos no dominio, ela basicamente é uma modelagem pronta para recebermos os dados de fora. Recebendo os dados de fora, podemos processar esses dados e quando tivermos os dados prontos, podemos pegar esses dados e jogar eles na nossa camada de dominio ou de negócio.

## usecase
Executa o fluxo de operações com as regras de negócio.
Define/Gera o fluxo de aplicação.
Como por exemplo: cria um banco, cria uma transação associada ao banco, cria uma chave, associa a chave a conta, etc. 
Esse fluxo um atrás do outro, é feito sempre aqui. Tudo integrado com o banco de dados, dominio, etc.

## cmd
Comandos registrados para iniciarmos a aplicação e seus serviços(CLI).
Responsável pelos comandos para iniciarmos a aplicação. Quando quisermos inicar a aplicação, essa pasta vai gerar a intercface do CMD. Por exemplo, colocamos "Code pix", "code pix kafka", "code pix health".

## domain / model
Coração e complexidade da aplicação e suas regras de negócios.
- domain: aonde vai estar as nossas entidades, a nossa modelagem do nosso negócio. Ai existe um banco, uma conta bancaria, transacao. Elas tem que ter uma coesão muito grande para não ter um dado invalido a qualquer momento. Vamos trabalhar com um dominio da forma mais rica o possivel.
Então, regras de negócio da nossa aplicação, vai ficar dentro de nosso dominio.

## infraestructure
Layer mais baixo nivel,  e nossa parte do repositorio. 
	- db: ajuda a fazer conexões como banco de dados. vamos utilizar um ORM simples, chamado goRM
	- repository: padrão que Ajuda a persistir e trazer dados do banco de dados para esses dados parar direto na camada de dominio. Independente banco de dados, arquivo txt. Normalmente, são chamados pelos "usecases"
	

# Recursos que vamos utilizar
- Docker
- Golang
- Apache Kafka
- Postgres