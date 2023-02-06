package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import java.util.Optional;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
// @RequestMapping(value="/vacationDaysLefts")
@Transactional
public class VacationDaysLeftController {

    @Autowired
    VacationDaysLeftRepository vacationDaysLeftRepository;

    @RequestMapping(
        value = "vacationDaysLefts/{id}/add",
        method = RequestMethod.PUT,
        produces = "application/json;charset=UTF-8"
    )
    public VacationDaysLeft add(
        @PathVariable(value = "id") String id,
        @RequestBody AddCommand addCommand,
        HttpServletRequest request,
        HttpServletResponse response
    ) throws Exception {
        System.out.println("##### /vacationDaysLeft/add  called #####");
        Optional<VacationDaysLeft> optionalVacationDaysLeft = vacationDaysLeftRepository.findById(
            id
        );

        optionalVacationDaysLeft.orElseThrow(() ->
            new Exception("No Entity Found")
        );
        VacationDaysLeft vacationDaysLeft = optionalVacationDaysLeft.get();
        vacationDaysLeft.add(addCommand);

        vacationDaysLeftRepository.save(vacationDaysLeft);
        return vacationDaysLeft;
    }

    @RequestMapping(
        value = "vacationDaysLefts/{id}/use",
        method = RequestMethod.PUT,
        produces = "application/json;charset=UTF-8"
    )
    public VacationDaysLeft use(
        @PathVariable(value = "id") String id,
        @RequestBody UseCommand useCommand,
        HttpServletRequest request,
        HttpServletResponse response
    ) throws Exception {
        System.out.println("##### /vacationDaysLeft/use  called #####");
        Optional<VacationDaysLeft> optionalVacationDaysLeft = vacationDaysLeftRepository.findById(
            id
        );

        optionalVacationDaysLeft.orElseThrow(() ->
            new Exception("No Entity Found")
        );
        VacationDaysLeft vacationDaysLeft = optionalVacationDaysLeft.get();
        vacationDaysLeft.use(useCommand);

        vacationDaysLeftRepository.save(vacationDaysLeft);
        return vacationDaysLeft;
    }
}
