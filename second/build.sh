if [[ -d build ]]; then
    rm -rf build
fi

GEN_STUB_SCRIPT=$(find /usr/lib/php -name gen_stub.php | head -n1) frankenphp extension-init main.go
CGO_ENABLED=1 \
    CGO_CFLAGS="-D_GNU_SOURCE $(php-config --includes) -Ibuild" \
    CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
    xcaddy \
        build \
        --output frankenphp-server \
        --with github.com/dunglas/frankenphp/caddy \
        --with example.com/myextension=. \
        && echo OK