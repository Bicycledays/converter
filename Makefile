build:
	cd cmd && go build main.go && chmod +x main && mv main converter

build-sartorius:
	cd cmd && go build main.go && chmod +x main && cp main ~/PhpstormProjects/sartorius/back/go/converter/