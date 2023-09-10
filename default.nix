{pkgs ? import <nixpkgs> {}}: pkgs.mkShell {
  buildInputs = with pkgs; [
    bombardier
    go
    gopls
    graphviz
    just
    unixtools.watch
  ];
}
