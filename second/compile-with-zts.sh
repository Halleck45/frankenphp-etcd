set -e

if [[ ! -d /tmp/php-src ]]; then
  git clone --branch=PHP-8.4 https://github.com/php/php-src.git /tmp/php-src
fi

cd /tmp/php-src
./buildconf --force
./configure \
		--enable-embed \
		--enable-zts \
		--disable-zend-signals \
		--enable-zend-max-execution-timers \
		--enable-debug

make -j"$(nproc)"
make install
ldconfig /etc/ld.so.conf.d
php --version