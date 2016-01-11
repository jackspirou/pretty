build:
	go install

release:
	go get github.com/mitchellh/gox
	go get github.com/jackspirou/tarpack
	gox
	tarpack pretty_*
