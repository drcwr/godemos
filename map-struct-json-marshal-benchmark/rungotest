echo "go test"
go test -bench=. -cpuprofile=cpu.prof
echo "pprof"
pprof -http=:6060 cpu.prof

#http://127.0.0.1:6060/ui/flamegraph?si=cpu