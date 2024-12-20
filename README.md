### ArchShell

ArchShell is minimized Linux wrapped in the EFI x86-64 application.

## Build dependencies:

- make utility
- docker with buildx plugin
- internet connection

## Build:

```
make clean
make
```

## Entry script source priority:

- **URI Device Path from current UEFI LoadOption.** The entry script URI will be determined as the end cut off the source URI delimited by `#`. The entry script file will be downloaded to the `/archshell/entrypoint` destination and then executed at the `/` working directory. If the source URI does not contain `#`, a command-line shell will be provided. If the script URI is determined and the file cannot be downloaded, a command-line shell will not be provided for security reasons.
- **BOOTP bootfile, DHCPv4 bootfile-name(67), DHCPv6 bootfile-uri(59) options when DHCP is allowed in IPv4 Device Path or IPv6 Device Path at current UEFI LoadOption.** The one who is the first received. The entry script URI will be taken, downloaded and executed in the same way as for the `URI Device Path from current UEFI LoadOption`. If the required options are not received via DHCP or the script file cannot be downloaded, the shell will not be provided for security reasons. But it can be taken when the source URI does not contain the delimiter.
- **Hard Drive Media Device Path supplemented with File Path Media Device Path from current UEFI LoadOption.** The entry script will be determined as the file located in the same directory as `archshell.efi` named `archshell.sh`. The target file will be symlinked to the `/archshell/entrypoint` and then executed at the `/` working directory. If the entry file is missing, a command-line shell will be provided.

## Bugs

- EFI FilePath separation is not supported (see https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#file-path-media-device-path)
- IPv6 network stack and VLAN network layer are not tested at all
