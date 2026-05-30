
{
  description = "slotchmod changes file permission with a slot";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    # 複数のシステム(Linux, Macなど)に簡単に対応するためのライブラリ
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "slotchmod";
          version = "1.1.2";
          src = ./.;
          vendorHash = "sha256-TP75WnspFTzcRrq8eUpbQGi5F/pIeYoCDMon13YmrJ8=";
        };

        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.go_1_26
            pkgs.gopls
          ];

          shellHook = ''
            echo "go development environment was activated"
          '';
        };
      }
    );
}
