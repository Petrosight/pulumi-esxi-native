# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'BootFirmwareType',
    'DiskType',
]


class BootFirmwareType(str, Enum):
    BIOS = "bios"
    EFI = "efi"


class DiskType(str, Enum):
    THIN = "thin"
    ZEROED_THICK = "zeroedthick"
    EAGER_ZEROED_THICK = "eagerzeroedthick"
    UNKNOWN = "Unknown"
