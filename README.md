# golang-project-structure
Golang project structure template(s)

# Idea
It would be nice to have a structure that would fit all type of projects regardless 
how complicated the project is and regardless of what architectural style it uses 
(Hexagonal, n-tier, ...)

Is it possible or do we need to change the project structure always completely based on what we are 
building?

My hypothesis is that it's possible to have a basic general template and then leave out
out some directories or add nested directories based on the concrete project

# Type of projects
- Backend service with business logic with RPC or REST HTTP API
- API service with limited amount of business logic, more mapping and validation
- Library providing go methods as an API 
- Binary
- Web project serving HTML, JS, Images

# Architecture style
- Hexagonal
- DDD - Domain Driven Design (might be close to hexagonal)
- 3-tier - Presentation, Business, Persistence
- Multilayer - Presentation, Application, Business, Persistence
- MVC - Model, View, Controller
- ... more ?


# Resources
- https://github.com/golang-standards/project-layout
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://rakyll.org/style-packages/
- https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html
- https://peter.bourgon.org/go-best-practices-2016/#repository-structure
- https://vsupalov.com/go-folder-structure/
- https://en.wikipedia.org/wiki/List_of_software_architecture_styles_and_patterns


# More ideas
- Sub-packages are not supported in Golang - what are the benefits, what are the downsides?

