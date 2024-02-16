bazel build :all

# Create the bazel directory if it doesn't exist
mkdir -p bazel

# Move the generated files to the bazel directory
mv bazel-bin bazel
mv bazel-out bazel
mv bazel-silence-of-the-lambdas bazel
mv bazel-testlogs bazel
