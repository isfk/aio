SHELL_FOLDER=$(
    # shellcheck disable=SC2164
    cd "$(dirname "$0")"
    pwd
)

# shellcheck disable=SC2045
for dirName in $(ls "$SHELL_FOLDER"/../proto/); do
    echo "${dirName}"
    protoc \
    -I=$SHELL_FOLDER/../proto \
    -I=$GOPATH/src \
    -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
    --micro_out=$SHELL_FOLDER/../proto --gogofaster_out=$SHELL_FOLDER/../proto \
    $SHELL_FOLDER/../proto/$dirName/*.proto
done