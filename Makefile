build:
	cc -o $(T) $(T).c

run:
	docker run --rm -v /Users/tomoya/9cc:/9cc -w /9cc compilerbook cc -o $(T) $(T).s
	docker run --rm -v /Users/tomoya/9cc:/9cc -w /9cc compilerbook ./$(T)
	docker run --rm -v /Users/tomoya/9cc:/9cc -w /9cc compilerbook echo $?