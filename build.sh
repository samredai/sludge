docker build --build-arg SLURM_TAG="slurm-22-05-3-1" -t build-sludge .
docker run -it -v ${PWD}:/root/sludge build-sludge
