default: run

gson-2.8.6.jar:
	wget -q https://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.6/gson-2.8.6.jar

Main.class: gson-2.8.6.jar Main.java
	@echo Compiling...
	@javac -cp ./gson-2.8.6.jar Main.java

run: gson-2.8.6.jar Main.class
	@echo Running...
	@java -cp ./gson-2.8.6.jar Main.java

clean:
	rm -f *.class gson-2.8.6.jar