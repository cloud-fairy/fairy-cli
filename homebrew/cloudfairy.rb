# class Fairy < Formula
#   desc "cloudfairy cli"
#   homepage "https://github.com/cloud-fairy/fairy-cli"
#   sha256 :no_check
#   license "MIT"
#   version: ""
  

#   on_macos do
#     url "https://github.com/cloud-fairy/fairy-cli/releases/download//fairy-darwin-amd64"
#   end

#   on_linux do
#     url "https://github.com/cloud-fairy/fairy-cli/releases/download//fairy-linux-amd64"
#   end

#   def install
#     bin.install "fairy"
#   end
# end

# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Fairy < Formula
  desc "cloudfairy cli"
  homepage "https://github.com/cloud-fairy/fairy-cli"
  sha256 :no_check
  url "https://github.com/cloud-fairy/fairy-cli/archive/refs/tags/v0.0.2.tar.gz"
  sha256 :no_check
  license "MIT"
  version: ""

  depends_on "go" => :build

  def install
    system "go", "build", "-o" "fairy", *std_go_args(ldflags: "-s -w")
    bin.install "fairy"
  end
end

