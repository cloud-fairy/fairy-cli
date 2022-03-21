class Fairy < Formula
  desc "cloudfairy cli"
  homepage "https://github.com/cloud-fairy/fairy-cli"
  sha256 :no_check
  license "MIT"
  version: "v0.0.1"
  

  on_macos do
    url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-darwin-amd64"
  end

  on_linux do
    url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-linux-amd64"
  end

  def install
    bin.install "fairy"
  end
end
