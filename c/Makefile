# Define the compiler
CC = clang

# Define compiler flags and max debug level
CFLAGS = -Wall -Wextra -Werror -g -fstandalone-debug -O0

# Define the target executable
TARGET = switcher_c

# Define the source files
SRCS = switcher.c display_info.c

# Define the object files
OBJS = $(SRCS:.c=.o)

# Default rule: build the target executable
all: $(TARGET)

# Rule to build the target executable
$(TARGET): $(OBJS)
	@echo "Linking $(TARGET) with CoreGraphics framework..."
	$(CC) $(OBJS) -o $(TARGET) -framework CoreGraphics

# Rule to compile source files into object files
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# Clean up the build files
clean:
	rm -f $(OBJS) $(TARGET)

# Phony targets
.PHONY: all clean

