# coding: utf-8


require 'pry'

PROJECT_HOME    = File.expand_path(File.dirname(__FILE__))
PROJECT_NAME    = "cloudbrowser"
BUILD_DIR       = "#{PROJECT_HOME}/build"
CONFIG_DIR      = "#{PROJECT_HOME}/configs"
ASSETS_DIR      = "#{PROJECT_HOME}/assets"
GOOS            = "$(go env GOOS)"
GOARCH          = "$(go env GOARCH)"

PACK_DIR        = "#{PROJECT_HOME}/pack"

task :default => %w(assets:statik build)

desc "production build #{PROJECT_NAME}"
task :build do
  build_target_dir = "#{PROJECT_HOME}/cmd"
  Dir.each_child(build_target_dir) do |s|
    sh <<-SHELL
      go build \
        -o "#{BUILD_DIR}/#{GOOS}_#{GOARCH}/#{s}" \
        #{build_target_dir}/#{s}
    SHELL
    puts "comple ok! :: #{s}"
  end
end

namespace :assets do
  desc "production build #{PROJECT_NAME}"
  task :statik do
    sh <<-SHELL
      statik -f -dest #{PACK_DIR} -p #{File.basename(CONFIG_DIR)} -src #{CONFIG_DIR} -include=*.yaml
      statik -f -dest #{PACK_DIR} -p #{File.basename(ASSETS_DIR)} -src #{ASSETS_DIR}
    SHELL
  end
end

namespace :install do
  desc "install dependencies"
  task :tools do
    sh <<-SHELL
      go get github.com/rakyll/statik
    SHELL
  end
end
