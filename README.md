# Replicating streaming issue with dagger v0.11.0

Run: `dagger run --progress plain go run ./main.go`

Expected: the dagger/container/services stderr is streamed as it generated.

Actual Behavoir: the dagger/container/services stderr is is dumped at the end of the run.

This is a regression from dagger v0.10.X
