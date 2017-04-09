TARGET_DIR = target

# Build
all:debug

debug:
	@if [ ! -d "$(TARGET_DIR)" ] ; then \
	mkdir $(TARGET_DIR); \
	fi
	go build -o $(TARGET_DIR)/tileset_$@

# Build with flags "-w -s"
release:
	@if [ ! -d "$(TARGET_DIR)" ] ; then \
	mkdir $(TARGET_DIR); \
	fi
	go build -ldflags "-w -s" -o $(TARGET_DIR)/tileset

test:
	go test

install:
	go install

clean:
	rm $(TARGET_DIR) -rf

# Add svn ignore
svn:
  chmod +x svn_init.sh
  ./svn_init.sh
