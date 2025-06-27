set -e
DIR=$(dirname "$0")
cd "$DIR"

if [[ -d build ]]; then
    rm -rf build
fi

GEN_STUB_SCRIPT=$(find /usr/local/lib/php -name gen_stub.php | head -n1) frankenphp extension-init myextension.go

cp go.mod go.sum build/
cd build

# issue in myextension.go: comment the 'import "runtime/cgo"' line
sed -i 's/import "runtime\/cgo"//g' myextension.go

# issue in generated code: replace all "ext_module_entry" by "myextension_module_entry"
sed -i 's/ext_module_entry/myextension_module_entry/g' myextension.go
sed -i 's/ext_module_entry/myextension_module_entry/g' myextension.c
sed -i 's/ext_module_entry/myextension_module_entry/g' myextension.h

CGO_ENABLED=1 \
    CGO_CFLAGS="-D_GNU_SOURCE -DZTS -DPHP_ZTS $(php-config --includes) -I." \
    CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
    xcaddy \
        build \
        --output frankenphp-server \
        --with github.com/dunglas/frankenphp/caddy \
        --with example.com/myextension=. \
        && echo OK