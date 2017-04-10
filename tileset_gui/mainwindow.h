#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QMessageBox>
#include <QFileDialog>
#include <QProcess>
#include <QUrl>
#include <QDesktopServices>
#include <QDebug>

namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();

private slots:
    void on_pushButton_start_clicked();

    void on_actionFile_triggered();

    void on_actionFolder_triggered();

    void on_actionAbout_triggered();

    void on_actionVisit_git_responsity_triggered();

    void on_actionExit_triggered();

    void on_pushButton_Origin_File_clicked();

    void on_pushButton_Origin_Folder_clicked();

    void on_pushButton_Out_File_clicked();

    void on_pushButton_Out_Folder_clicked();

    void on_output();

    void on_finished();

private:
    Ui::MainWindow *ui;
    QString last_Origin_Dir = ".";
    QString last_Out_Dir = ".";
    QProcess *process;
};

#endif // MAINWINDOW_H
