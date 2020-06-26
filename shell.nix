with import <nixpkgs> {}; {
  devEnv = stdenv.mkDerivation {
        name = "devgo";
	buildInputs = with pkgs; [
		stdenv
		gcc
		pkgconfig
		autoconf
		automake
		libtool
		gnumake
		glibc.static
		go
	];
	CFLAGS="-I${pkgs.glibc.dev}/include";
	LDFLAGS="-L${pkgs.glibc}/lib";
  };
}
