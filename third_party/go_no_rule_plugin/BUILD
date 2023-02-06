package(default_visibility = ["PUBLIC"])

go_module(
    name = "pgx",
    install = ["..."],
    module = "github.com/jackc/pgx/v5",
    version = "v5.2.0",
    deps = [
        ":pgpassfile",
        ":pgservicefile",
        ":puddle",
        ":xcrypto",
        ":xtext",
    ],
)

go_module(
    name = "pgpassfile",
    module = "github.com/jackc/pgpassfile",
    version = "v1.0.0",
)

go_module(
    name = "puddle",
    module = "github.com/jackc/puddle/v2",
    version = "v2.1.2",
    deps = [
        ":atomic",
        ":xsync",
    ],
)

go_module(
    name = "xsync",
    install = ["semaphore"],
    module = "golang.org/x/sync",
    version = "v0.1.0",
)

go_module(
    name = "atomic",
    module = "go.uber.org/atomic",
    version = "v1.10.0",
)

go_module(
    name = "pgservicefile",
    module = "github.com/jackc/pgservicefile",
    version = "v0.0.0-20200714003250-2b9c44734f2b",
)

go_module(
    name = "xcrypto",
    install = ["pbkdf2"],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20220829220503-c86fa9a7ed90",
)

go_module(
    name = "xtext",
    install = [
        "transform",
        "cases",
        "language",
        "runes",
        "unicode/norm",
        "width",
        "secure/bidirule",
        "secure/precis",
        "internal",
        "internal/language",
        "internal/tag",
        "internal/language/compact",
        "unicode/bidi",
    ],
    module = "golang.org/x/text",
    version = "v0.6.0",
)
