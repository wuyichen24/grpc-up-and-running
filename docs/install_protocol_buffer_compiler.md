## Install Protocol Buffer Compiler
### For Ubuntu/Debian
```bash
$ sudo apt-get install autoconf automake libtool curl make g++ unzi
```
### For Other Platforms
- **Step 1**: Download the source code.
   - **Way 1**: Download from Github.
      - Use this link to download: [https://github.com/protocolbuffers/protobuf/releases/latest](https://github.com/protocolbuffers/protobuf/releases/latest)
      - Pick the right version

        | Language | Pick This File |
        |---|---|
        | C++ | protobuf-cpp-[VERSION].(tar.gz\|zip) |
        | Java | protobuf-java-[VERSION].(tar.gz\|zip) |
        | C# | protobuf-csharp-[VERSION].(tar.gz\|zip) |
        | JavaScript | protobuf-csharp-[VERSION].(tar.gz\|zip) |
        | Objective-C | protobuf-objectivec-[VERSION].(tar.gz\|zip) |
        | PHP | protobuf-php-[VERSION].(tar.gz\|zip) |
        | Python | protobuf-python-[VERSION].(tar.gz\|zip) |
        | Ruby | protobuf-ruby-[VERSION].(tar.gz\|zip) |
        | All Languages | protobuf-all-[VERSION].(tar.gz\|zip) |
   - **Way 2**: Clone from Github.
     ```bash
      git clone https://github.com/protocolbuffers/protobuf.git
      cd protobuf
      git submodule update --init --recursive
      ./autogen.sh
     ```
- **Step 2**: Install the source code.
  ```bash
  ./configure
  make
  make check
  sudo make install
  ```

