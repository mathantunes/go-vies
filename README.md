# go-vies

VIES API golang

Exposing a simple API to handle VAT validation

```go
import (
    "github.com/mathantunes/go-vies"

    "fmt"
)
    func Using() error {
        v := vies.NewVIES(nil) // specify a different endpoint, otherwise it will utilize the default
        resp, err := v.Validate("FI25160553")
        // Failures are described in err
        if err != nil {
            return err;
        }
        // VAT Validation check can be found on
        if !resp.Valid {
            fmt.Errorf("It seems like the VAT provided is not valid :(")
        }
    }
```

## Up Next

* Grpc implementation for running it on docker or container orchestration software
* Lambda implementation with SAM template
