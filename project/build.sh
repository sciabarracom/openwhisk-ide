wsk property set \
    --apihost 'http://openwhisk:3233' \
    --auth '23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP'
zip /tmp/main.zip main.ts tsconfig.ts
wsk action update hellots /tmp/main.zip --docker actionloop/action-typescript-v3.7
wsk action invoke hellots -r
mkdir .theia
wsk action invoke hellots -p debugWith vscode -r
echo '{"version":"0.2.0", "configurations": [] }' |\
 jq  --argfile data <(wsk action invoke hellots -p debugWith vscode -r)  \
 '.configurations = [$data]' >.theia/launch.json

