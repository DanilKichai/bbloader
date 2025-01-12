# syntax=docker/dockerfile:1

FROM --platform=linux/amd64 golang:latest as builder
  COPY . /usr/src/archshell
  RUN \
    CGO_ENABLED="0" \
      go build \
        -C /usr/src/archshell/cmd/bootstrap \
        -ldflags='-extldflags=-static' \
        -o /usr/local/bin/bootstrap \
        .

FROM --platform=linux/amd64 archlinux:base as wrapper
  ADD hack/etc/pacman.d/mirrorlist /etc/pacman.d/

  RUN \
    pacman \
      --sync \
      --sysupgrade \
      --refresh \
      --noconfirm \
      --needed \
      core/binutils \
      core/curl \
      core/dosfstools \
      core/e2fsprogs \
      core/efibootmgr \
      core/less \
      core/lvm2 \
      core/nano \
      extra/arch-install-scripts \
      extra/debootstrap \
      extra/kexec-tools \
      extra/polkit \
      extra/tcpdump \
      extra/traceroute

  WORKDIR /target
  RUN \
    tar \
      --create \
      --file - \
      --one-file-system \
      --exclude=target \
      --exclude=var/cache \
      --exclude=run/* \
      / | tar --extract --file -
  ADD build/package/target/ /target/
  COPY --from=builder /usr/local/bin/bootstrap /target/usr/local/bin/
  RUN echo -e "debug\ndebug" | passwd --root /target

  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      core/linux
  ADD build/package/toolchain/ /
  RUN mkinitcpio --preset wrapper

FROM scratch as target
  COPY --from=wrapper /boot/wrapper.efi archshell.efi
