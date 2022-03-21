class FairyCli < Formula
  desc "cloudfairy cli"
  homepage "https://github.com/cloud-fairy/fairy-cli"
  version "v0.0.1"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-windows-amd64.exe"
    when /darwin|mac os/
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-darwin-amd64"
    when /linux/
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-linux-amd64"
    when /solaris|bsd/
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-freebsd-amd64"
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-windows-386.exe"
    when /darwin|mac os/
      :unknown
    when /linux/
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-linux-386"
    when /solaris|bsd/
      url "https://github.com/cloud-fairy/fairy-cli/releases/download/v0.0.1/fairy-freebsd-386"
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "fairy"
  end

end
