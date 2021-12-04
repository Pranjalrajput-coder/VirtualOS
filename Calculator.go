package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)



func calculatorapp() {
		
		
		output:=""
		output_panel :=widget.NewLabel(output)

		isHistory := false
		historyStr:=""
		history:=widget.NewLabel(historyStr)		// for just showing the result above the output panel

		var historyArr [] string;
		
		historybtn :=widget.NewButton("History",func(){
			if isHistory{
				historyStr = ""
			}else{
				for i:=len(historyArr)-1; i>=0; i--{
					historyStr = historyStr + historyArr[i]
					historyStr +="\n"
				}
			}
			isHistory = !isHistory;
			history.SetText(historyStr)
		})

		backbtn :=widget.NewButton("Back",func(){
			if len(output)>0 {
			output=output[:len(output)-1]
			output_panel.SetText(output)
			}
		})
		
		clearbtn :=widget.NewButton("Clear",func(){
			output = ""
			output_panel.SetText(output)
		})
		
		bracketleftbtn :=widget.NewButton("(",func(){
				output = output + "("
			output_panel.SetText(output)
		})
		
		bracketrightbtn :=widget.NewButton(")",func(){
				output = output+")"
			output_panel.SetText(output)
		})
		
		dividebtn :=widget.NewButton("/",func(){
				output = output+"/"
			output_panel.SetText(output)
		})
		
		sevenbtn :=widget.NewButton("7",func(){
				output = output+"7"
			output_panel.SetText(output)
		})
		
		eightbtn :=widget.NewButton("8",func(){
				output = output+"8"
			output_panel.SetText(output)
		})
		
		ninebtn :=widget.NewButton("9",func(){
				output = output+"9"
			output_panel.SetText(output)
		})
		
		multiplybtn :=widget.NewButton("*",func(){
				output = output+"*"
			output_panel.SetText(output)
		})
		
		fourbtn :=widget.NewButton("4",func(){
				output = output+"4"
			output_panel.SetText(output)
		})
		
		fivebtn :=widget.NewButton("5",func(){
				output = output+"5"
			output_panel.SetText(output)
		})
		
		sixbtn :=widget.NewButton("6",func(){
				output = output+"6"
			output_panel.SetText(output)
		})
		
		minusbtn :=widget.NewButton("-",func(){
				output = output+"-"
			output_panel.SetText(output)
		})
		
		onebtn :=widget.NewButton("1",func(){
				output = output+"1"
			output_panel.SetText(output)
		})
		
		twobtn :=widget.NewButton("2",func(){
				output = output+"2"
			output_panel.SetText(output)
		})
		
		threebtn :=widget.NewButton("3",func(){
				output = output+"3"
			output_panel.SetText(output)
		})
		
		plusbtn :=widget.NewButton("+",func(){
				output = output+"+"
			output_panel.SetText(output)
		})
		
		zerobtn :=widget.NewButton("0",func(){
				output = output+"0"
			output_panel.SetText(output)
		})
		
		dotbtn :=widget.NewButton(".",func(){
				output = output+"."
			output_panel.SetText(output)
		})
		
		equalbtn :=widget.NewButton("=",func(){
			expression, err := govaluate.NewEvaluableExpression(output);
			if err == nil{
				result, err := expression.Evaluate(nil);
				if err==nil{
					answer := strconv.FormatFloat(result.(float64),'f',-1,64);		// for converting the string in int format
					strToAppend:=output +" = " +answer;
					historyArr= append(historyArr,strToAppend);
					output=answer;
					}else{
					output="Error";
				}
			}else{
				output="Error";
			}
			output_panel.SetText(output)
		})

		calculatorContainer := container.NewVBox(
			container.NewVBox(
			output_panel,
			history,
			container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historybtn,
				backbtn,
				),
				container.NewGridWithColumns(4,
				clearbtn,
				bracketleftbtn,
				bracketrightbtn,
				dividebtn,
			),
			container.NewGridWithColumns(4,
				ninebtn,
				eightbtn,
				sevenbtn,
				multiplybtn,
			),
			container.NewGridWithColumns(4,
				sixbtn,
				fivebtn,
				fourbtn,
				minusbtn,
				),
			container.NewGridWithColumns(4,
				onebtn,
				twobtn,
				threebtn,
				plusbtn,
				),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
				zerobtn,
				dotbtn,),
				equalbtn,
			),),),
		)

		w:=myApp.NewWindow("Calculator");
		w.Resize(fyne.NewSize(500,280));
	w.CenterOnScreen()

		w.SetContent(container.NewBorder(DeskBtn,nil,nil,nil,calculatorContainer))
		w.Show();
}