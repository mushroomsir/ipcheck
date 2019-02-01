function test {
    go test ./... -coverprofile="coverage.out" 
    go tool cover -func="coverage.out"
    Remove-Item "coverage.out"
}

switch ($args[0]) {
    "test" { test }
    Default { "Not support " + $args[0]  }
}