import os, sys

# Open a file
fd = os.open( "/etc/test_break", os.O_RDWR|os.O_CREAT )

# Write one string
os.write(fd, "This is test")

# Now you can use fsync() method.
# Infact here you would not be able to see its effect.
os.fsync(fd)

# Now read this file from the beginning
os.lseek(fd, 0, 0)
str = os.read(fd, 100)
print "Read String is : ", str

# Close opened file
os.close( fd )
