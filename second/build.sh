set -e
DIR=$(dirname "$0")
cd "$DIR"

if [[ -d build ]]; then
    rm -rf build
fi

GEN_STUB_SCRIPT=$(find /usr/local/lib/php -name gen_stub.php | head -n1) frankenphp extension-init myextension.go convert.go

cp go.mod go.sum build/
cd build
CGO_ENABLED=1 \
    CGO_CFLAGS="-D_GNU_SOURCE -DZTS -DPHP_ZTS $(php-config --includes) -Ibuild" \
    CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
    xcaddy \
        build \
        --output frankenphp-server \
        --with github.com/dunglas/frankenphp/caddy \
        --with example.com/myextension=. \
        && echo OK