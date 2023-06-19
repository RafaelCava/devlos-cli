<div align="center" id="top"> 
  <img src="https://cdn.worldvectorlogo.com/logos/golang-gopher.svg" alt="Devlos Cli" style="width: 100px"/>
</div>


<h1 align="center">Devlos Cli</h1>

<p align="center">
  <img alt="Principal linguagem do projeto" src="https://img.shields.io/github/languages/top/RafaelCava/devlos-cli?color=56BEB8">

  <img alt="Quantidade de linguagens utilizadas" src="https://img.shields.io/github/languages/count/RafaelCava/devlos-cli?color=56BEB8">

  <img alt="Tamanho do repositório" src="https://img.shields.io/github/repo-size/RafaelCava/devlos-cli?color=56BEB8">

  <img alt="Licença" src="https://img.shields.io/github/license/RafaelCava/devlos-cli?color=56BEB8">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/RafaelCava/devlos-cli?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/RafaelCava/devlos-cli?color=56BEB8" /> -->

  <img alt="Github stars" src="https://img.shields.io/github/stars/RafaelCava/devlos-cli?color=56BEB8" />
</p>

<!-- Status -->

<h4 align="center"> 
	🚧  Devlos Cli 🚀 Em construção...  🚧
</h4> 

<hr>

<p align="center">
  <a href="#dart-sobre">Sobre</a> &#xa0; | &#xa0; 
  <a href="#sparkles-funcionalidades">Funcionalidades</a> &#xa0; | &#xa0;
  <a href="#rocket-tecnologias">Tecnologias</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-pré-requisitos">Pré requisitos</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-começando">Começando</a> &#xa0; | &#xa0;
  <a href="#memo-licença">Licença</a> &#xa0; | &#xa0;
  <a href="https://github.com/RafaelCava" target="_blank">Autor</a>
</p>

<br>

## :dart: Sobre ##

Projeto criado para melhorar a experiência dos desenvolvedores, durante o desenvolvimento de novas features

## :sparkles: Funcionalidades ##

:heavy_check_mark: Funcionalidade 1: Inicia o ambiente de desenvolvimento\
:heavy_check_mark: Funcionalidade 2: Finaliza o ambiente de desenvolvimento\
:x: Funcionalidade 3: Busca dados de um ticket informado\
:x: Funcionalidade 4: Cria tasks no JIRA\
:x: Funcionalidade 5: Edita tasks no JIRA\
:x: Funcionalidade 6: Deleta tasks no JIRA
## :rocket: Tecnologias ##

As seguintes ferramentas foram usadas na construção do projeto:

- [Go](https://go.dev/)
- [Cobra-cli](https://github.com/spf13/cobra)
- [Promptui](https://github.com/manifoldco/promptui)
- [Survey](https://github.com/AlecAivazis/survey)

## :white_check_mark: Pré requisitos ##

Antes de começar :checkered_flag:, você precisa ter o [Git](https://git-scm.com) e o [Go](https://go.dev/) ou [Docker](https://www.docker.com/) instalados em sua maquina.

## :checkered_flag: Começando ##

```bash
# Clone este repositório
$ git clone https://github.com/RafaelCava/devlos-cli

# Entre na pasta
$ cd devlos-cli

# Configure no seu arquivo .zshrc ou .bashrc a variável de ambiente PATH_TALLOS apontando para o diretório onde esta os repositórios da tallos.
# É muito importante não mudar o nome dos diretórios da tallos para manter o padrão atual, a cli esta seguindo esse padrão e não fará nada além dele.

# Build a CLI
$ GOFLAGS=-mod=mod go build -o /bin/devlos-cli main.go

# Caso deseje utilizar o docker para buildar a CLI
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.19 go build -buildvcs=false 
$ sudo mv ./devlos-cli /bin/
```

## :memo: Licença ##

Este projeto está sob licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.


Feito com :heart: por <a href="https://github.com/RafaelCava" target="_blank">Rafael Cavalcante</a>

Deixe a sua :star:

&#xa0;

<a href="#top">Voltar para o topo</a>
