{
  description = "Flake for developing friendly";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { nixpkgs
    , flake-utils
    , ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShell = pkgs.mkShell {

          packages = with pkgs; [
            go
            nushell
            goose
          ];

          shellHook = ''
            exec nu
          '';
        };
      }
    );
}
