import subprocess

proc = subprocess.Popen(
    [".\\main.exe"],
    stdin=subprocess.PIPE,
    stdout=subprocess.PIPE,
    stderr=subprocess.PIPE,
    text=True
)

stdout, stderr = proc.communicate(input='{"numbers":[1,2,3,4,5]}')

print("STDOUT:", stdout)
print("STDERR:", stderr)
print("Return code:", proc.returncode)