#!/bin/bash

FEATURE_FILE="gpm.features"

OPTIONS=(
  1 "Packages over SSH" off
  2 "Packages over HTTPS (Default)" on
  3 "Gpm-Autobuild (Default)" on
  4 "Gpm-API (optional, for automation only)" off
)

CHOICES=$(dialog --checklist "Select GPM build features" 15 60 4 \
  "${OPTIONS[@]}" 3>&1 1>&2 2>&3)

clear
echo "# GPM Feature Config" > $FEATURE_FILE
[[ $CHOICES =~ 1 ]] && echo "USE_SSH=true" >> $FEATURE_FILE || echo "USE_SSH=false" >> $FEATURE_FILE
[[ $CHOICES =~ 2 ]] && echo "USE_HTTPS=true" >> $FEATURE_FILE || echo "USE_HTTPS=false" >> $FEATURE_FILE
[[ $CHOICES =~ 3 ]] && echo "ENABLE_AUTOBUILD=true" >> $FEATURE_FILE || echo "ENABLE_AUTOBUILD=false" >> $FEATURE_FILE
[[ $CHOICES =~ 4 ]] && echo "ENABLE_API=true" >> $FEATURE_FILE || echo "ENABLE_API=false" >> $FEATURE_FILE

echo "Feature selections saved to $FEATURE_FILE"
