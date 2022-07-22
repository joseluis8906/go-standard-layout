# `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

Following the above recommendations, this directory is structure as follows:

* `/internal/application`: this directory is composed of `commands` and `queries` directories, following the Bertrand Meyer's `command query separation` concept. Some command examples could be `CreateOrder`, `AddItem`, and some queries like `GetItems`, `GetTaxes`. [`Martin Fowler's article`](https://www.martinfowler.com/bliki/CommandQuerySeparation.html)
* `/internal/domain`: this directory must be used to host the `boundaries` covered by the service, for example `Product`, `Order`, `Item`, `Tax`. All of this boundaries are composed by `Entities`, `Value Objects` and `Aggregates` [`Domain Models and Bounded Contexts`](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/architect-microservice-container-applications/identify-microservice-domain-model-boundaries) 
* `/internal/infrastructure`: this is the place you will use to host infrastructure code like repository implementations and http/grpc external service clients. Must be divided into protocols, for example `postgres`, `http`, `grpc`, `mongodb`.
