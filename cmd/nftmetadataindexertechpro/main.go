// cmd/nftmetadataindexertechpro/main.go
package main

import (
"flag"
"log"
"os"

"nftmetadataindexertechpro/internal/nftmetadataindexertechpro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmetadataindexertechpro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
