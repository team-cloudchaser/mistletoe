[CmdletBinding()]
param (
    [Parameter(Mandatory=$false)]
    [String]
    $Config,
    [Parameter(Mandatory=$false)]
    [String]
    $Target
)

if(!$Config) {
    Write-Output "Base config not specified. Available configs:"
    (Get-ChildItem -Filter *.json .\data).BaseName | ForEach-Object { Write-Output "- $($_)"}
    Write-Output "`nUse ""./load <config> <target>"" to connect to a server."
    exit 1;
}

if(!(Test-Path -PathType Leaf ("$(Join-Path -Path "data" -ChildPath $Config).json"))){
    Write-Host "Base Config ""$Config"" does not exist"
    exit 1;
}

if(!$Target){
    Write-Output "Target server not specified. Available targets:"
    (Get-Content .\data\remote.tsv.bin) | Select-Object -Skip 1 | ForEach-Object { Write-Output "- $(($_ -split "`t")[0])" }
    Write-Output "`nUse ""./load.ps1 $Config <target>"" to connect to a server."
    exit 1;
}

$MISTLETOE_LINE = ((((Get-Content .\data\remote.tsv.bin) | Select-String "^$($Target)") -split "`t") | Select-Object -Skip 1)
if(!$MISTLETOE_LINE){
    Write-Output "Specified target does not exist."
    exit 1
}

$tmpFile="tmp.json"
$lineArgs = $MISTLETOE_LINE
if (Get-Command -Name xray -ErrorAction SilentlyContinue) {
    Remove-Item $tmpFile *> $null;
    Get-Content -LiteralPath "$(Join-Path -Path "data" -ChildPath $Config).json" > $tmpFile
    ((Get-Content -LiteralPath $tmpFile) -replace "__SERVER__|__SID__|__PBK__|__UUID__", { switch ($_.Value) {
        "__SERVER__" { $lineArgs[0] }
        "__SID__" { $lineArgs[1] }
        "__PBK__" { $lineArgs[2] }
        "__UUID__" { $lineArgs[3] }
    }}) | Set-Content -Path $tmpFile
    xray run -c "$tmpFile"
} else {
    Write-Output "Required proxy toolchain is not installed."
}
exit