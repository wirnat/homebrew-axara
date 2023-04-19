class Axara < Formula
  desc "CLI for Code Generator"
  homepage "https://github.com/wirnat/axara"
  url "https://github.com/wirnat/axara/archive/refs/tags/v1.2.0.tar.gz" # Ganti URL dengan URL tautan unduhan package Anda
  sha256 "6d7b66d824347d48cf92ad9006cc7b6e1c04aca9dd780af00dfb336214938fff" # Ganti HASH_SHA256_FILE dengan hash SHA256 dari file unduhan Anda
  version "1.0.0"

  def install
    bin.install "axara" # Ganti "nama-package" dengan nama binary file package Anda
  end
end
