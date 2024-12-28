{
  description = "A todo list manager for git repositiories";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
    goBin = pkgs.buildGoModule {
      name = "git-todo";
      src = ./.;
      vendorHash = "sha256-eKeUhS2puz6ALb+cQKl7+DGvm9Cl+miZAHX0imf9wdg=";

      CGO_ENABLED=0;
    };

  in {
    packages.x86_64-linux.default = goBin;
  };
}
