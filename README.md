# big boo
#### Largest boo ever attempted
---

### Building
```bash
# Windows 
gradle.bat clean build

# nix 
gradlew clean build 
```

### Squashing commits
```bash
git rebase -i HEAD~${#_of commits}

git push origin ${branch} --force with lease
```