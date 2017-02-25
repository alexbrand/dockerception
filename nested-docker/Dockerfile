FROM centos:7

RUN yum -y update && yum -y clean all

RUN yum -y remove docker docker-common container-selinux docker-selinux

RUN yum install -y yum-utils && yum-config-manager \
    --add-repo \
    https://docs.docker.com/engine/installation/linux/repo_files/centos/docker.repo

RUN yum makecache fast && yum -y install docker-engine

# Enable systemd
RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == \
systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

VOLUME [ "/sys/fs/cgroup" ]

RUN systemctl enable docker

CMD ["/usr/sbin/init"]
