# Golang project structure


TODO check this: https://medium.com/@kujtimii.h/creating-a-todo-app-using-gokit-cli-20f066a58e1

Idea: Separate directories for logical packages and layered packages (todo find correct names)
E.g. application directory with layers(handler, service, repositories) and next to it directory with some packages that 
doesnt fit to layers, but are somehow logical - http, security, ....


### The idea
It would be nice to have a structure that would fit all type of projects regardless 
how complicated the project is and regardless of what architectural style it uses 
(Hexagonal, n-tier, ...)

Is it possible or do we need to change the project structure always completely based on what we are 
building?

## Hypotheses 1
It's possible to have a basic general template and then leave out
out some directories or add nested directories based on the concrete project

## Hypotheses 2 - Counter Hypotheses for Hypotheses 1
There are too many different type of projects and it's important how big the project is to select
the appropriate directory structure.

# Approach
1. Create projects categorization
1. Collects examples of existing directory structure from real projects
1. Try to combine them and create so called "master" directory structure that would combine all
1. ....
1. Profit :)

## TODO thoughts:
- business applications vs tools vs things like kubernetes (what type of software is that?), https server, proxy, 
do you write these things the same way?


# Categorization of projects
## Type of projects
TODO refine those categories ... think about projects like smaller cli tools, or bigger like docker , ...

- Backend service with business logic with RPC or REST HTTP API
- API service with limited amount of business logic, more mapping and validation
- Library providing go methods as an API 
- Binary
- Web project serving HTML, JS, Images

## Size of the projects
- Small library providing a couple of methods
- A simple micro-service which provides key - value store via a REST API, has persistence
- Relatively small micro-service that has 1 Aggregate, Persistence, and communicates with an external service
- Middle size service that has 1 aggregate, persistence, REST API and RPC API and communicates with 
2 other services via RPC
- ... ?

## Architecture styles
- Unnamed - just create packages with names that make sense
- N-tier (Presentation, Application, Business, Persistence)
- Hexagonal
- MVC - Model, View, Controller
- ... 

# Example directory structures

## 1) Simple Lib
    .
    ├── example              # Directory with examples
    ├── gobreaker.go         # The library
    ├── gobreaker_test.go    
    ├── go.mod
    ├── go.sum
    └── README.md
    
## 2) Simple service, TODO: does it make sense to show such simple example? ... or is it useless? :/
    .
    ├── main.go         # TODO: Is one file example really "real world production ready thing?"
    ├── go.mod
    ├── go.sum
    └── README.md

## 3) Service with "no architecture" just name packages properly
    .
    ├── api                  # Proto/Swagger contract definition     
    ├── cmd                  
    ├── plugin                  
    ├── k8s                  
    ├── proxy                  
    ├── stats                  
    ├── template                  
    ├── handler              # RPC handlers (kinda controllers)  
    ├── go.mod
    ├── go.sum       
    ├── Jenkinsfile
    ├── main.go              # Main file of the service
    ├── Makefile
    ├── LICENSE
    └── README.md

## 4) Commandline tool
... TODO
    
## 5) N-Tier
    .
    ├── contract              # Contract definitions
    │   ├── proto             
    │   └── swagger          
    ├── doc                   # Documentation (images, charts, more md files if needed)
    ├── k8s                   # K8s deployment files
    ├── pkg                   # We might use "internal" instead
    │   ├── application   
    │   ├── config   
    │   ├── controller    
    │   ├── entities   
    │   └── repository
    ├── .gitignore
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum    
    ├── Jenkinsfile
    ├── main.go               # Main file of the service
    ├── Makefile
    ├── LICENSE
    └── README.md
    
## 6) Hexagonal
You will probably not have more aggregates in one service, but it might make sense in some situations.
For example at the beginning of the project where splitting into more services would be slowing down the development.
Other example is when 

    .
    ├── build                       # Compiled files (not commited, but appears in the project after first build)
    ├── contract                    # 
    │   ├── proto                   # Proto contract definition
    │   └── swagger                 # Swagger contract definition
    ├── doc                         # Documentation (images, charts, more md files if needed)
    ├── gradle                      # Gradle wrapper
    ├── k8s                         # K8s deployment files
    ├── pkg or internal?            # No code from the service is not supposed to be used as a library, but do we need to user the "internal" directory there?
    │   ├── config  
    │   ├── application             # Application layer, Responsibility: orcherstration of models, persistence, transactions, auth    
    │   ├── client                  # Domain servicse (DDD lingua)
    │   ├── handler                 # RPC / REST "ports", basically controllers from N-tier
    │   ├── domain                  # The domain model (sometimes called "model")
    │   └── repository              # Persistence "ports"
    ├── .gitignore
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum    
    ├── Jenkinsfile
    ├── main.go                       # Main file of the service
    ├── Makefile
    ├── LICENSE
    └── README.md


    
## 7) ??
.... 


# Results (in progress)

## Hypothesis 2 confirmation
- This repository https://github.com/golang-standards/project-layout has quite some many (8000 atm) which 
suggests that many people like it.


# Resources
- https://github.com/golang-standards/project-layout
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://rakyll.org/style-packages/
- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2
- https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html
- https://peter.bourgon.org/go-best-practices-2016/#repository-structure
- https://vsupalov.com/go-folder-structure/
- https://en.wikipedia.org/wiki/List_of_software_architecture_styles_and_patterns