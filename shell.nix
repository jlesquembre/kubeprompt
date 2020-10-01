with import <nixpkgs> { };
pkgs.mkShell {
  buildInputs = [
    gotools
    gopls
  ];
}
