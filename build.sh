set -e
echo 'Start build astools binary executable files...'
version=$(go version)
if [[ "$version" =~ "go version" ]]; then
  echo "$version"
else
  echo "Error: please install Go first."
  exit 4
fi
GOOS=android GOARCH=arm64 go build -o astools_android_arm64 cmd/cli/main.go
GOOS=darwin GOARCH=arm64 go build -o astools_darwin_arm64 cmd/cli/main.go
GOOS=darwin GOARCH=amd64 go build -o astools_darwin_amd64 cmd/cli/main.go
GOOS=linux GOARCH=amd64 go build -o astools_linux_amd64 cmd/cli/main.go
GOOS=linux GOARCH=arm64 go build -o astools_linux_arm64 cmd/cli/main.go
GOOS=windows GOARCH=arm64 go build -o astools_windows_arm64.exe cmd/cli/main.go
GOOS=windows GOARCH=amd64 go build -o astools_windows_amd64.exe cmd/cli/main.go
echo "build done."
ls -larth ./astools*
