<a name="readme-top" id="readme-top"></a>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#how-to-use-wizworld">How to use</a></li>
        <li><a href="#available-commands">Available Commands</a></li>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li>
      <a href="#testing">Testing</a>
      <ul>
        <li><a href="#unit-testing">Unit Testing</a></li>
      </ul>
    </li>
    <li>
      <a href="#project-structure">Project Structure</a>
      <ul>
        <li><a href="#actors">Actors</a></li>
      </ul>
    </li>
        <li>
      <a href="#assumptions-notes-and-todos">Assumptions, notes and todos</a>
    </li>
  </ol>
</details>

<br/>

<!-- ABOUT THE PROJECT -->

# About The Project

WizWorld Cli - A command line tool to help you navigate through the world of wizards

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<br/>

# Getting Started

## How to use wizworld

In order to run wizworld, first ou need to build the binary. For building the binary you have multiple options:

1. Use make: `make build`
2. Use golang: `go build -o <your-destination-path/wizworld> cmd/main.go`

To run the program you can use the following options:

```bash
  ./build/wizworld
```

or

```bash
  cd build
  ./wizworld
```

or if you use custom path

```bash
  cd <your-destination-path>
  ./wizworld
```

To see the available commands, use `./wizworld --help`

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Available commands

### ingredients
Ingredients command lists the available ingredients.

Example: ``` ./wizworld ingredients ```

To get more info about the command, use: ```./wizworld ingredients --help```

<br/>

### elixirs
Elixirs command lists elixirs and also you can search elixirs by ingredients. You can filter single or multiple ingredients.

Example: ``` ./wizworld elixirs -i "Neem oil, Jewelweed" ```

To get more info about the command, use: ```./wizworld elixirs --help```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<br/>

## Prerequisites
### Project Dependencies
| Name | Version | Description | Download & Docs |   |
|------|---------|-------------|------|---|
| Golang | < v1.19 | An open-source programming language supported by Google.  | [Link](https://go.dev/)  |   |

### Project Dependencies for development
| Name | Version | Description | Download & Docs |   |
|------|---------|-------------|------|---|
| Mockery | < v2.14 | mockery provides the ability to easily generate mocks for golang interfaces | [Link](https://github.com/vektra/mockery)  |   |

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<br/>

## Testing
### Unit testing
To run unit tests use ```make test``` command


<p align="right">(<a href="#readme-top">back to top</a>)</p>
<br/>

## Project Structure
The project follows module based, port-adapter architecture. A module contains different providers (etc: service, command). Each provider has its own functionality and purpose. Each domain module has a root file. The root file is responsible for constructing  providers, exporting functionality and managing dependencies.

Third party packages:
- cobra-cli
- color
- table

Third party packages for testing:
- gofakeit/v6
- testify


### Actors
Driving (or primary) actors are the ones that initiate the interaction. For example, a driving adapter could be a controller which is the one that takes the (user) input and passes it to the Application via a Port or Cli.

Driven (or secondary) actors are the ones that are “kicked into behavior” by the Application. For example, a database Adapter is called by the Application so that it fetches a certain data set from persistence.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<br/>

## Assumptions, notes and todos

- The Ingredient query parameter of the https://wizard-world-api.herokuapp.com/Elixirs endpoint is only works with a single ingredient. In order to filter by multiple ingredients, I added additional logic to the code (service.go)
- TODO: I tested the majority of the code, but I still have some missing cases, like io and output operation errors
- TODO: It would be nice to add system/e2e tests to the code

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<br/>
