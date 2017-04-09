#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    process = new QProcess();
    connect(process,SIGNAL(readyRead()),this,SLOT(on_output()));
    ui->setupUi(this);
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::on_pushButton_start_clicked()
{
#ifdef Q_OS_WIN32
    QString cmd = "./tilesetCore.exe";
#else
    QString cmd = "./tilesetCore";
#endif
    QString useGpu = ui->checkBox_useGpu->checkState()?"gpu":"cpu";

    QString input = ui->lineEdit_Origin->text();
    QString output = ui->lineEdit_Out->text();

    QStringList params;

    params << useGpu << input << output;

    process->start(cmd,params);
    ui->pushButton_start->setEnabled(false);
    process->waitForFinished();
    ui->pushButton_start->setEnabled(true);

}

bool MainWindow::on_output(){
    ui->output->append(process->readAll());
}

void MainWindow::on_actionFile_triggered()
{
    on_pushButton_Origin_File_clicked();
}

void MainWindow::on_actionFolder_triggered()
{
    on_pushButton_Origin_Folder_clicked();
}

void MainWindow::on_actionAbout_triggered()
{
    QMessageBox::information(this,"关于","xp/va素材转mv by garfeng\n放大算法：waifu2x");
}

void MainWindow::on_actionVisit_git_responsity_triggered()
{
    QUrl url("https://github.com/garfeng/tileset");
    QDesktopServices::openUrl(url);
}

void MainWindow::on_actionExit_triggered()
{
    this->close();
}

void MainWindow::on_pushButton_Origin_File_clicked()
{
    //char imgpath[128];
    QString imgPath = QFileDialog::getOpenFileName(this,"选择图片",last_Origin_Dir, tr("Image Files(*.png)"));
    //qDebug()<< imgPath;
    if(imgPath != ""){
        ui->lineEdit_Origin->setText(imgPath);
        last_Origin_Dir = imgPath;
    }
}

void MainWindow::on_pushButton_Origin_Folder_clicked()
{
   QString imgDir = QFileDialog::getExistingDirectory(this,"选择文件夹",last_Origin_Dir);

   if(imgDir!=""){
       ui->lineEdit_Origin->setText(imgDir);
       last_Origin_Dir = imgDir;
   }
   qDebug() << imgDir;
}

void MainWindow::on_pushButton_Out_File_clicked()
{
    QString imgPath = QFileDialog::getSaveFileName(this,"选择输出文件",last_Out_Dir,tr("Image Files(*.png)"));
    if(imgPath!=""){
        ui->lineEdit_Out->setText(imgPath);
        last_Out_Dir = imgPath;
    }
}

void MainWindow::on_pushButton_Out_Folder_clicked()
{
    QString imgDir = QFileDialog::getExistingDirectory(this,"选择输出目录",last_Out_Dir);
    if(imgDir != "") {
        ui->lineEdit_Out->setText(imgDir);
        last_Out_Dir= imgDir;
    }
}
