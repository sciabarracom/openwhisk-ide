wsk property set --apihost 'http://openwhisk:3233' --auth '23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP'
wsk action create hellots main.ts --docker actionloop/action-typescript-v3.7
wsk action invoke hellots -r
wsk action invoke hellots -p debugWith vscode -r