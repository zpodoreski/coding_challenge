#!/bin/bash

echo
echo "Deploying..."
echo

# get input
package=$1

# copy pkg to destination
echo
echo "Moving pkg to /$package"
echo
cp -R pkg functions/"$package"

# rename imports
echo
echo "Renaming imports to point to local pkg folder in $package"
echo
find functions/"$package" -type f -name '*.go' -print0 | xargs -0 sed -i '' -e "s?coding_challenge\/pkg?coding_challenge\/functions\/$package\/pkg?"


# run goimports on all files in destination
echo
echo "Running go goimports on package $package"
echo
find functions/"$package" -type f -name \*.go -exec goimports -w {} \;

cd functions/"$package" || exit

echo
echo "Running go mod init and tidy on package $package"
echo
go mod init
go mod tidy


echo
echo "Deploying $package to GCP http functions"
echo
gcloud functions deploy "$package" --entry-point "$package" --runtime go116 --trigger-http --allow-unauthenticated


echo
echo "Doing cleanup after deploy"
echo
rm -rf pkg
rm -f go.mod
rm -f go.sum

echo
echo "Renaming imports to initial state"
echo
find . -type f -name '*.go' -print0 | xargs -0 sed -i '' -e "s?coding_challenge\/functions\/$package\/pkg?coding_challenge\/pkg?"

echo
echo "Running goimports"
echo
find . -type f -name \*.go -exec goimports -w {} \;

echo
echo "DONE!!!"
echo
